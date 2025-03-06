package engine

import (
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func EnvPutPath(key string, value ...string) {
	src := os.Getenv(key)
	if len(src) != 0 {
		return
	}
	fn.Panic(os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)))
}
func EnvVar(key string) string {
	return os.Getenv(key)
}
func EvalFile(e *Engine, path string) any {
	return fn.Panic1(e.RunCode(CompileFile(EnvExpand(path), true))).Export()
}
func EvalFiles(e *Engine, paths ...string) (r []any) {
	for _, s := range paths {
		r = append(r, fn.Panic1(e.RunCode(CompileFile(EnvExpand(s), true))).Export())
	}
	return
}
func Mkdir(path string) {
	pt := EnvExpand(path)
	if FileExists(pt) {
		return
	}
	fn.Panic(os.Mkdir(pt, os.ModePerm))
}
func MkdirAll(path string) {
	pt := EnvExpand(path)
	if FileExists(pt) {
		return
	}
	fn.Panic(os.MkdirAll(pt, os.ModePerm))
}

func WriteBinaryFile(path string, data []byte) {
	fn.Panic(os.WriteFile(EnvExpand(path), data, os.ModePerm))
}
func WriteTextFile(path string, data string) {
	fn.Panic(os.WriteFile(EnvExpand(path), []byte(data), os.ModePerm))
}
func ReadBinaryFile(path string) []byte {
	return fn.Panic1(os.ReadFile(EnvExpand(path)))
}
func ReadTextFile(path string) string {
	return string(fn.Panic1(os.ReadFile(EnvExpand(path))))
}
func Chdir(path string) error {
	return os.Chdir(EnvExpand(path))
}
func EnvPut(key string, value ...string) error {
	src := os.Getenv(key)
	if len(src) != 0 {
		return nil
	}
	return os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string]))
}
func EnvSetPath(key string, value ...string) error {
	tar := fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)
	return os.Setenv(key, tar)
}
func EnvSet(key string, value ...string) error {
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	return os.Setenv(key, tar)
}
func EnvPrepend(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = tar + pathListSeparator + src
	}
	return os.Setenv(key, tar)
}
func EnvAppendPath(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)
	if len(src) != 0 {
		tar = src + pathListSeparator + tar
	}
	return os.Setenv(key, tar)
}
func EnvAppend(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = src + pathListSeparator + tar
	}
	return os.Setenv(key, tar)
}
func EnvPrependPath(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)
	if len(src) != 0 {
		tar = tar + pathListSeparator + src
	}
	return os.Setenv(key, tar)
}

func Pwd() (string, error) {
	return os.Getwd()
}
func FileExists(path string) bool {
	_, err := os.Stat(EnvExpand(path))
	return err == nil || !os.IsNotExist(err)
}
func EnvExpand(path string) string {
	if strings.HasPrefix(path, "@") {
		return fixPath(os.ExpandEnv(filepath.Join(root, strings.TrimPrefix(path, "@"))))
	}
	if strings.HasPrefix(path, ".") {
		p, e := filepath.Abs(path)
		if e != nil {
			return fixPath(os.ExpandEnv(path))
		}
		return fixPath(os.ExpandEnv(p))
	}
	return fixPath(os.ExpandEnv(path))

}
func Stat(path string) map[string]any {
	entry, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return map[string]any{
		//"dir":      entry.IsDir(),
		"name":     entry.Name(),
		"mode":     entry.Mode().String(),
		"size":     entry.Size(),
		"modified": entry.ModTime().Format("2006-01-02 15:04:05.000Z07"),
	}
}
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && !fi.IsDir()
}
func IsSymlink(path string) bool {
	fi, err := os.Lstat(path)
	return err == nil && fi.Mode()&os.ModeSymlink != 0
}
func FileCopy(src, dest string, force *bool) error {
	f := force != nil && *force
	src = EnvExpand(src)
	dest = EnvExpand(src)
	if IsDir(src) {
		return copyDirectory(src, dest, f)
	}
	return copyFile(src, dest, f)
}
func FileMove(src, dest string, force *bool) error {
	f := force != nil && *force
	src = EnvExpand(src)
	dest = EnvExpand(dest)
	err := os.Rename(src, dest)
	if err == nil {
		return nil
	}
	// 如果错误是跨设备链接错误（EXDEV），执行复制+删除
	var linkError *os.LinkError
	if errors.As(err, &linkError) {
		var errno syscall.Errno
		if errors.As(linkError.Err, &errno) && errors.Is(errno, syscall.EXDEV) {
			// 检查目标是否存在
			if _, err := os.Stat(dest); !os.IsNotExist(err) {
				if f {
					if err := os.RemoveAll(dest); err != nil {
						return fmt.Errorf("无法删除目标目录: %v", err)
					}
				} else {
					return fmt.Errorf("目标目录已存在，使用 -f 强制覆盖")
				}
			}

			// 执行跨设备复制
			if err := copyDirectory(src, dest, f); err != nil {
				return err
			}

			// 删除原始目录
			return os.RemoveAll(src)
		}
	}

	return fmt.Errorf("move %s to %s failure: %v", src, dest, err)
}

// 递归复制目录
func copyDirectory(src, dest string, force bool) error {
	if !force && FileExists(dest) {
		return fmt.Errorf("%s exists", dest)
	}
	// 创建目标目录
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return err
	}

	// 遍历源目录
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dest, relPath)

		// 处理目录
		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// 处理文件
		return copyFile(path, targetPath, true)
	})
}

// 复制单个文件
func copyFile(src, dest string, force bool) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	if !force && FileExists(dest) {
		return fmt.Errorf("%s exists", dest)
	}
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}
	// 保持文件权限
	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dest, si.Mode())
}

func CreateSymlink(src, dest string) error {
	// 统一路径格式
	src = filepath.Clean(src)
	dest = filepath.Clean(dest)

	// 检查源路径是否存在
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("source not exists: %s", src)
	}

	// 处理不同操作系统的符号链接创建
	switch runtime.GOOS {
	case "windows":
		return createWindowsSymlink(src, dest)
	default: // linux/darwin 等类 Unix 系统
		return os.Symlink(src, dest)
	}
}

// Windows 专用符号链接创建
func createWindowsSymlink(src, dest string) error {
	// 获取源文件信息
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("获取源信息失败: %v", err)
	}

	// 确定符号链接类型
	var flags uint32
	if srcInfo.IsDir() {
		flags = syscall.SYMBOLIC_LINK_FLAG_DIRECTORY
	}

	// 转换路径为NT格式
	srcNT, err := syscall.UTF16PtrFromString(src)
	if err != nil {
		return err
	}
	destNT, err := syscall.UTF16PtrFromString(dest)
	if err != nil {
		return err
	}

	// 使用Windows API创建符号链接
	err = syscall.CreateSymbolicLink(destNT, srcNT, flags)
	if err != nil {
		// 处理常见错误
		switch {
		case errors.Is(err, syscall.ERROR_PRIVILEGE_NOT_HELD):
			return fmt.Errorf("privilege required")
		case errors.Is(err, syscall.ERROR_ALREADY_EXISTS):
			return fmt.Errorf("destnation exists: %s", dest)
		default:
			return fmt.Errorf("create fail: %v", err)
		}
	}
	return nil
}

func ReadSymlink(path string) (string, error) {
	path = EnvExpand(path)
	target, err := os.Readlink(path)
	if err != nil {
		return "", fmt.Errorf("read fail: %v", err)
	}
	return EnvExpand(target), nil
}
