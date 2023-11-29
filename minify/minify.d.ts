declare module 'go/minify' {
    /**
     * @param v ECMAScript version
     */
    export function setESVersion(v: 0 | 2014 | 2015 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022)

    // @ts-ignore
    import {Reader, Writer} from 'go/io'

    export interface Minify {

        bytes(mime: Mime, src: Uint8Array): Uint8Array

        string(mime: Mime, src: string): string

        reader(mime: Mime, src: Reader): Reader

        minify(mime: Mime, out: Writer, src: Reader)

    }

    type Mime =
        'text/css'
        | 'text/html'
        | 'image/svg+xml'
        | 'application/javascript'
        | 'application/json'
        | 'application/xml'

    export function minifier(): Minify

    /**
     * guess and minify file
     * @param input file to minified
     * @param output file to write output
     */
    export function fileToFile(input, output: string)
    /**
     *
     * @param input file to minified
     * @param output file to write output
     * @param mime the mime type
     */
    export function fileToFile(input, output: string, mime: Mime)
    /**
     *
     * @param input file to minified
     * @param output folder to write output with input file name
     */
    export function fileToFolder(input, output: string)
    /**
     *
     * @param input file to minified
     * @param output folder to write output with input file name
     * @param mime the mime type
     */
    export function fileToFolder(input, output: string, mime: Mime)
    /**
     *
     * @param input folder to scan files for minified
     * @param output folder to write output with input file name
     */
    export function folderToFolder(input, output: string)
    /**
     *
     * @param input folder to scan files for minified
     * @param output folder to write output with input file name
     * @param mime the mime type
     */
    export function folderToFolder(input, output: string, mime: Mime)
}