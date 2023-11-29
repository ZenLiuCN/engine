declare module "go/esbuild" {
    type StderrColor = 0 | 1 | 2
    type LogLevel = 0 | 1 | 2 | 3 | 4 | 5
    type Charset = 0 | 1 | 2
    type TreeShaking = 0 | 1 | 2
    /**
     * Include,Exclude
     */
    type SourcesContent = 0 | 1
    /**
     * Default,None,Inline,EndOfFile,Linked,External
     */
    type LegalComments = 0 | 1 | 2 | 3 | 4 | 5
    /**
     * Transform, Preserve, Automatic
     */
    type JSX = 0 | 1 | 2
    type Drop = 1 | 2
    /**
     *    Default,ESNext,ES5,ES2015,ES2016,ES2017,ES2018,ES2019,ES2020,ES2021,ES2022
     */
    type Target = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10
    /**
     *    None,Inline,Linked,External,InlineAndExternal
     */
    type SourceMap = 0 | 1 | 2 | 3 | 4
    /**
     * LoaderNone , LoaderBase64, LoaderBinary, LoaderCopy, LoaderCSS, LoaderDataURL, LoaderDefault, LoaderEmpty, LoaderFile, LoaderGlobalCSS, LoaderJS, LoaderJSON, LoaderJSX, LoaderLocalCSS, LoaderText, LoaderTS, LoaderTSX
     */
    type Loader = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16
    type MangleQuoted = 0 | 1
    /**
     * Default,Browser,Node,Neutral
     */
    type Platform = 0 | 1 | 2 | 3
    /**
     * default,IIFE,CommonJs,ESM
     */
    type Format = 0 | 1 | 2 | 3
    /**
     * default,external
     */
    type Packages = 0 | 1
    /**
     *    Chrome,Deno,Edge,Firefox,Hermes,IE,IOS,Node,Opera,Rhino,Safari
     */
    type EngineName = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10

    export interface Engine {
        name: EngineName
        version: string
    }

    export interface TransformOptions {
        color: StderrColor // Documentation: https://esbuild.github.io/api/#color
        logLevel: LogLevel // Documentation: https://esbuild.github.io/api/#log-level
        logLimit: number // Documentation: https://esbuild.github.io/api/#log-limit
        logOverride: { [key: string]: LogLevel } // Documentation: https://esbuild.github.io/api/#log-override

        sourcemap: SourceMap // Documentation: https://esbuild.github.io/api/#sourcemap
        sourceRoot: string // Documentation: https://esbuild.github.io/api/#source-root
        sourcesContent: SourcesContent // Documentation: https://esbuild.github.io/api/#sources-content

        target: Target // Documentation: https://esbuild.github.io/api/#target
        Engines: Engine[]        // Documentation: https://esbuild.github.io/api/#target
        Supported: { [key: string]: boolean } // Documentation: https://esbuild.github.io/api/#supported

        platform: Platform // Documentation: https://esbuild.github.io/api/#platform
        format: Format // Documentation: https://esbuild.github.io/api/#format
        globalName: string // Documentation: https://esbuild.github.io/api/#global-name

        mangleProps: string // Documentation: https://esbuild.github.io/api/#mangle-props
        reserveProps: string // Documentation: https://esbuild.github.io/api/#mangle-props
        mangleQuoted: MangleQuoted // Documentation: https://esbuild.github.io/api/#mangle-props
        MangleCache: { [key: string]: any }// Documentation: https://esbuild.github.io/api/#mangle-props
        drop: Drop // Documentation: https://esbuild.github.io/api/#drop
        dropLabels: string[]               // Documentation: https://esbuild.github.io/api/#drop-labels
        minifyWhitespace: boolean // Documentation: https://esbuild.github.io/api/#minify
        minifyIdentifiers: boolean // Documentation: https://esbuild.github.io/api/#minify
        minifySyntax: boolean // Documentation: https://esbuild.github.io/api/#minify
        lineLimit: number //int // Documentation: https://esbuild.github.io/api/#line-limit
        charset: Charset // Documentation: https://esbuild.github.io/api/#charset
        treeShaking: TreeShaking // Documentation: https://esbuild.github.io/api/#tree-shaking
        ignoreAnnotations: boolean // Documentation: https://esbuild.github.io/api/#ignore-annotations
        legalComments: LegalComments // Documentation: https://esbuild.github.io/api/#legal-comments

        jSX: JSX // Documentation: https://esbuild.github.io/api/#jsx
        jSXFactory: string // Documentation: https://esbuild.github.io/api/#jsx-factory
        jSXFragment: string // Documentation: https://esbuild.github.io/api/#jsx-fragment
        jSXImportSource: string // Documentation: https://esbuild.github.io/api/#jsx-import-source
        jSXDev: boolean // Documentation: https://esbuild.github.io/api/#jsx-dev
        jSXSideEffects: boolean // Documentation: https://esbuild.github.io/api/#jsx-side-effects

        tsconfigRaw: string // Documentation: https://esbuild.github.io/api/#tsconfig-raw
        banner: string // Documentation: https://esbuild.github.io/api/#banner
        footer: string // Documentation: https://esbuild.github.io/api/#footer

        define: Record<string, string>// Documentation: https://esbuild.github.io/api/#define
        pure: string[] // Documentation: https://esbuild.github.io/api/#pure
        keepNames: boolean // Documentation: https://esbuild.github.io/api/#keep-names

        sourcefile: string // Documentation: https://esbuild.github.io/api/#sourcefile
        loader: Loader // Documentation: https://esbuild.github.io/api/#loader
    }

    export interface TransformResult {
        errors?: Message[]
        warnings?: Message[]

        code?: Uint8Array
        map?: Uint8Array
        legalComments?: Uint8Array

        mangleCache: Record<string, any>
    }

    export function transform(input: string, options: TransformOptions): TransformResult

    export interface StdinOptions {
        contents: string
        resolveDir: string
        sourcefile: string
        loader: Loader
    }

    export interface EntryPoint {
        inputPath: string
        outputPath: string
    }

    export interface BuildResult {
        errors: Message[]
        warnings: Message[]

        outputFiles: OutputFile[]
        metafile: string
        mangleCache: Record<string, any>
    }

    export interface OutputFile {
        path: string
        contents: Uint8Array
        hash: string
    }

    export interface BuildOptions {
        color: StderrColor // Documentation: https://esbuild.github.io/api/#color
        logLevel: LogLevel // Documentation: https://esbuild.github.io/api/#log-level
        logLimit: number // Documentation: https://esbuild.github.io/api/#log-limit
        logOverride: Record<string, LogLevel> // Documentation: https://esbuild.github.io/api/#log-override

        sourcemap: SourceMap // Documentation: https://esbuild.github.io/api/#sourcemap
        sourceRoot: string // Documentation: https://esbuild.github.io/api/#source-root
        sourcesContent: SourcesContent // Documentation: https://esbuild.github.io/api/#sources-content

        target: Target // Documentation: https://esbuild.github.io/api/#target
        engines: Engine[]        // Documentation: https://esbuild.github.io/api/#target
        supported: Record<string, boolean> // Documentation: https://esbuild.github.io/api/#supported

        mangleProps: string // Documentation: https://esbuild.github.io/api/#mangle-props
        reserveProps: string // Documentation: https://esbuild.github.io/api/#mangle-props
        mangleQuoted: MangleQuoted // Documentation: https://esbuild.github.io/api/#mangle-props
        mangleCache: Record<string, any> // Documentation: https://esbuild.github.io/api/#mangle-props
        drop: Drop // Documentation: https://esbuild.github.io/api/#drop
        dropLabels: string[]               // Documentation: https://esbuild.github.io/api/#drop-labels
        minifyWhitespace: boolean // Documentation: https://esbuild.github.io/api/#minify
        minifyIdentifiers: boolean // Documentation: https://esbuild.github.io/api/#minify
        minifySyntax: boolean // Documentation: https://esbuild.github.io/api/#minify
        lineLimit: number // Documentation: https://esbuild.github.io/api/#line-limit
        charset: Charset // Documentation: https://esbuild.github.io/api/#charset
        treeShaking: TreeShaking // Documentation: https://esbuild.github.io/api/#tree-shaking
        ignoreAnnotations: boolean // Documentation: https://esbuild.github.io/api/#ignore-annotations
        legalComments: LegalComments // Documentation: https://esbuild.github.io/api/#legal-comments

        jSX: JSX // Documentation: https://esbuild.github.io/api/#jsx-mode
        jSXFactory: string // Documentation: https://esbuild.github.io/api/#jsx-factory
        jSXFragment: string // Documentation: https://esbuild.github.io/api/#jsx-fragment
        jSXImportSource: string // Documentation: https://esbuild.github.io/api/#jsx-import-source
        jSXDev: boolean // Documentation: https://esbuild.github.io/api/#jsx-dev
        jSXSideEffects: boolean // Documentation: https://esbuild.github.io/api/#jsx-side-effects

        define: Record<string, string> // Documentation: https://esbuild.github.io/api/#define
        pure: string[]          // Documentation: https://esbuild.github.io/api/#pure
        keepNames: boolean // Documentation: https://esbuild.github.io/api/#keep-names

        globalName: string // Documentation: https://esbuild.github.io/api/#global-name
        bundle: boolean // Documentation: https://esbuild.github.io/api/#bundle
        preserveSymlinks: boolean // Documentation: https://esbuild.github.io/api/#preserve-symlinks
        splitting: boolean // Documentation: https://esbuild.github.io/api/#splitting
        outfile: string // Documentation: https://esbuild.github.io/api/#outfile
        metafile: boolean // Documentation: https://esbuild.github.io/api/#metafile
        outdir: string // Documentation: https://esbuild.github.io/api/#outdir
        outbase: string // Documentation: https://esbuild.github.io/api/#outbase
        absWorkingDir: string // Documentation: https://esbuild.github.io/api/#working-directory
        platform: Platform // Documentation: https://esbuild.github.io/api/#platform
        format: Format // Documentation: https://esbuild.github.io/api/#format
        external: string[]          // Documentation: https://esbuild.github.io/api/#external
        packages: Packages // Documentation: https://esbuild.github.io/api/#packages
        alias: Record<string, string> // Documentation: https://esbuild.github.io/api/#alias
        mainFields: string[]          // Documentation: https://esbuild.github.io/api/#main-fields
        conditions: string[]          // Documentation: https://esbuild.github.io/api/#conditions
        loader: Record<string, Loader> // Documentation: https://esbuild.github.io/api/#loader
        resolveExtensions: string[]          // Documentation: https://esbuild.github.io/api/#resolve-extensions
        tsconfig: string // Documentation: https://esbuild.github.io/api/#tsconfig
        tsconfigRaw: string // Documentation: https://esbuild.github.io/api/#tsconfig-raw
        outExtension: Record<string, string> // Documentation: https://esbuild.github.io/api/#out-extension
        publicPath: string // Documentation: https://esbuild.github.io/api/#public-path
        inject: string[]          // Documentation: https://esbuild.github.io/api/#inject
        banner: Record<string, string> // Documentation: https://esbuild.github.io/api/#banner
        footer: Record<string, string> // Documentation: https://esbuild.github.io/api/#footer
        nodePaths: string[]          // Documentation: https://esbuild.github.io/api/#node-paths

        entryNames: string // Documentation: https://esbuild.github.io/api/#entry-names
        chunkNames: string // Documentation: https://esbuild.github.io/api/#chunk-names
        assetNames: string // Documentation: https://esbuild.github.io/api/#asset-names

        entryPoints: string[]     // Documentation: https://esbuild.github.io/api/#entry-points
        entryPointsAdvanced: EntryPoint[] // Documentation: https://esbuild.github.io/api/#entry-points

        stdin?: StdinOptions // Documentation: https://esbuild.github.io/api/#stdin
        write: boolean // Documentation: https://esbuild.github.io/api/#write
        allowOverwrite: boolean // Documentation: https://esbuild.github.io/api/#allow-overwrite
        plugins: Plugin[]      // Documentation: https://esbuild.github.io/plugins/
    }

    export function build(options: BuildOptions): BuildResult

    export interface ServeOnRequestArgs {
        remoteAddress: string
        method: string
        path: string
        status: number //int
        timeInMS: number //int // The time to generate the response, not to send it
    }

    export interface ServeOptions {
        port: number //uint16
        host: string
        servedir: string
        keyfile: string
        certfile: string
        fallback: string
        onRequest: (args: ServeOnRequestArgs) => void
    }

    export interface WatchOptions {
    }

    export interface ServeResult {
        port: string //uint16
        host: string
    }

    export interface ContextError {
        errors: Message[] // Option validation errors are returned here
    }

    export interface BuildContext {
        // Documentation: https://esbuild.github.io/api/#rebuild
        rebuild(): BuildResult

        // Documentation: https://esbuild.github.io/api/#watch
        watch(options: WatchOptions):Err

        // Documentation: https://esbuild.github.io/api/#serve
        serve(options: ServeOptions): {result:ServeResult,err:Err}

        cancel()

        dispose()
    }

    export function context(options: BuildOptions): { context: BuildContext, err: ContextError }

    export interface AnalyzeMetafileOptions {
        color: boolean
        verbose: boolean
    }

    export interface Location {
        file: string
        namespace: string
        line: number//int // 1-based
        column: number// int // 0-based, in bytes
        length: number// int // in bytes
        lineText: string
        suggestion: string
    }

    export interface Note {
        text: string
        location?: Location
    }

    export interface Message {
        id: string
        pluginName: string
        text: string
        location?: Location
        notes: Note[]

        // Optional user-specified data that is passed through unmodified. You can
        // use this to stash the original error, for example.
        Detail: any
    }

    export interface FormatMessagesOptions {
        terminalWidth: number//int
        kind: 0 | 1 // error | warning
        color: boolean
    }

    export function analyzeMetafile(metaFile: string, opts: AnalyzeMetafileOptions): string


    export function formatMessages(msg: Message[], opts: FormatMessagesOptions): string[]
}
