declare module 'go/chrome' {
    // @ts-ignore
    import {Err,Maybe,Stringer} from 'go'
    // @ts-ignore
    import {Writer,Closer} from 'go/io'
    // @ts-ignore
    import {Duration} from 'go/time'

    export interface TargetID extends Stringer{}
    export function toTargetID(id:string):TargetID
    export interface BrowserContextID extends Stringer{}
    export function toBrowserContextID(id:string):BrowserContextID
    export interface FrameID extends Stringer{

    }

    export function toFrameID(id:string):FrameID
    export interface TargetSessionID extends Stringer{

    }
    export function toTargetSessionID(id:string):TargetSessionID
    export interface NodeID {}
    export function toNodeID(id:number):NodeID
    export function fromNodeID(id:NodeID):number


    export interface Action {
    }

    export interface ValuedAction<T> extends Action {
        /**
         * the value,may not exist before submit to chrome
         */
        value(): T

    }
    export interface Browser{

    }
    export interface BrowserOption{

    }
    export function withDialTimeout(dur:Duration):BrowserOption

    export interface TargetInfo{
        targetID:TargetID
        type:string
        title:string
        url:string
        attached:boolean
        openerID?:TargetID
        canAccessOpener?:boolean
        openerFrameID?:FrameID
        browserContextID?:BrowserContextID
        subtype?:string
    }
    export class Chrome  extends Closer{
        constructor(execOptions:ExecOption[],...options: ContextOption[])
        /**
         *
         * @param url root websocket url, eg: ws://127.0.0.1:65533
         * @param remoteOptions
         * @param opts
         */
        constructor(url: string,remoteOptions?:RemoteOption[],...opts:ContextOption[])
        constructor(...options: ContextOption[])
        targets():Maybe<TargetInfo>
        submit(...act: Action[]): Err
        // createBrowser(url:string,...opt:BrowserOption[]):Maybe<Browser>

        shutdown()
    }
    export interface ContextOption{}
    export function withTargetID(targetId:TargetID):ContextOption
    export function withExistingBrowserContext(browserId:BrowserContextID):ContextOption
    export function withBrowserOption(...browserOptions:BrowserOption[]):ContextOption
    export interface RemoteOption{}

    //region exec options
    export interface ExecOption {
    }

    export function execPath(path: string): ExecOption

    export function flag(flag: string, val: string | boolean): ExecOption

    export function env(...pairs: string[]): ExecOption

    export function userDataDir(path: string): ExecOption

    export function proxyServer(proxy: string): ExecOption

    export function ignoreCertErrors(): ExecOption

    export function windowSize(width, height: number): ExecOption

    export function userAgent(userAgent: string): ExecOption

    export const noSandbox: ExecOption

    export const noFirstRun: ExecOption

    export const headless: ExecOption

    export const disableGPU: ExecOption

    export const noModifyURL: RemoteOption

    export function combinedOutput(writer: Writer): ExecOption

    export function wsUrlReadTimeout(duration: Duration): ExecOption
    //endregion

    export function navigate(url: string): Action

    export type NaviEntryId = number

    export interface NaviEntry {
        id: NaviEntryId
        url: string
        userTypedURL: string
        title: string
        transitionType: "link" | "typed" | "address_bar" | "auto_bookmark" | "auto_subframe" | "manual_subframe" |
            "generated" | "auto_toplevel" | "form_submit" | "reload" | "keyword" | "keyword_generated" | "other"
    }

    export interface NaviEntries {
        currentIndex: NaviEntryId
        entries: NaviEntry[]
    }

    export function navigate(url: string): Action

    export function navigateBack(): Action

    export function navigateForward(): Action

    export function reload(): Action

    export function stopLoading(): Action

    export function location(): ValuedAction<string>

    export function title(): ValuedAction<string>

    export function navigateToHistoryEntry(id: NaviEntryId): Action

    export function navigationEntries(): ValuedAction<NaviEntries>

    export interface PollOption {
    }

    export interface BackendNode {
        nodeType: number
        nodeName: string
        backendNodeID: number
    }

    /**
     * NodeTypeElement                = 1
     *
     * NodeTypeAttribute              = 2
     *
     * NodeTypeText                   = 3
     *
     * NodeTypeCDATA                  = 4
     *
     * NodeTypeEntityReference        = 5
     *
     * NodeTypeEntity                 = 6
     *
     * NodeTypeProcessingInstruction  = 7
     *
     * NodeTypeComment                = 8
     *
     * NodeTypeDocument               = 9
     *
     * NodeTypeDocumentType           = 10
     *
     * NodeTypeDocumentFragment       = 11
     *
     * NodeTypeNotation               = 12
     *
     */
    export type NodeType = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12
    /**
     * NodeReady = 128
     *
     * NodeVisible = 64
     *
     * NodeHighlighted = 32
     *
     */
    export type NodeState = 128 | 64 | 32
    export type PseudoType = "first-line" |
        "first-letter" |
        "before" |
        "after" |
        "marker" |
        "backdrop" |
        "selection" |
        "target-text" |
        "spelling-error" |
        "grammar-error" |
        "highlight" |
        "first-line-inherited" |
        "scrollbar" |
        "scrollbar-thumb" |
        "scrollbar-button" |
        "scrollbar-track" |
        "scrollbar-track-piece" |
        "scrollbar-corner" |
        "resizer" |
        "input-list-button" |
        "view-transition" |
        "view-transition-group" |
        "view-transition-image-pair" |
        "view-transition-old" |
        "view-transition-new"

    export interface DomNode {
        nodeID: number
        parentID: number
        backendNodeID: number
        nodeType: NodeType
        nodeName: string
        localName: string
        nodeValue: string
        childNodeCount?: number
        children?: DomNode[]
        attributes?: string[]
        documentURL?: string
        baseURL?: string
        publicID?: string
        systemID?: string
        internalSubset?: string
        xmlVersion?: string
        name?: string
        value?: string
        pseudoType?: PseudoType
        pseudoIdentifier?: string
        shadowRootType?: "user-agent" | "open" | "closed"
        frameID?: string
        contentDocument?: DomNode
        shadowRoots?: DomNode[]
        templateContent?: DomNode
        pseudoElements?: DomNode[]
        distributedNodes?: BackendNode[]
        isSVG?: boolean
        compatibilityMode?: "QuirksMode" | "LimitedQuirksMode" | "NoQuirksMode"
        assignedSlot: string

        partialXPathByID(): string

        attributeValue(name: string): string

        attribute(name: string): string //TODO

        partialXPath(): string

        partialXPath(): string

        fullXPathByID(): string

        fullXPath(): string

        dump(prefix, indent: string, nodeIDs: boolean): string
    }

    export function pollingInterval(dur: Duration): PollOption

    export function pollingTimeout(dur: Duration): PollOption

    export function pollingMutation(): PollOption

    export function pollingInFrame(frame: DomNode): PollOption

    export function pollingArgs(...args: any[]): PollOption

    export function poll(expr: string, res: any, ...options: PollOption[]): Action

    export function pollFunction(func: string, res: any, ...options: PollOption[]): Action

    //Query
    export interface QueryOption {
    }

    export const FromNode: QueryOption
    export const ByFunc: QueryOption
    export const ByQuery: QueryOption
    export const ByQueryAll: QueryOption
    export const ByID: QueryOption
    export const BySearch: QueryOption
    export const ByJSPath: QueryOption
    export const ByNodeID: QueryOption
    export const NodeReady: QueryOption
    export const NodeVisible: QueryOption
    export const NodeNotVisible: QueryOption
    export const NodeEnabled: QueryOption
    export const NodeSelected: QueryOption
    export const NodeNotPresent: QueryOption

    export function atLeast(n: number): QueryOption

    export function fromNode(n: DomNode): QueryOption

    export function retryInterval(n: Duration): QueryOption

    export function waitReady(sel: string, ...options: QueryOption[]): Action

    export function waitVisible(sel: string, ...options: QueryOption[]): Action

    export function waitNotVisible(sel: string, ...options: QueryOption[]): Action

    export function waitEnabled(sel: string, ...options: QueryOption[]): Action

    export function waitSelected(sel: string, ...options: QueryOption[]): Action

    export function waitNotPresent(sel: string, ...options: QueryOption[]): Action

    export function query(sel: string, ...options: QueryOption[]): Action

    export function focus(sel: string, ...options: QueryOption[]): Action

    export function blur(sel: string, ...options: QueryOption[]): Action

    export function clear(sel: string, ...options: QueryOption[]): Action

    export function click(sel: string, ...options: QueryOption[]): Action

    export function doubleClick(sel: string, ...options: QueryOption[]): Action

    export function submit(sel: string, ...options: QueryOption[]): Action

    export function reset(sel: string, ...options: QueryOption[]): Action

    export function scrollIntoView(sel: string, ...options: QueryOption[]): Action

    export function nodes(sel: string, ...options: QueryOption[]): ValuedAction<DomNode[]>

    export function nodeIDs(sel: string, ...options: QueryOption[]): ValuedAction<number[]>

    export function outerHTML(sel: string, ...options: QueryOption[]): ValuedAction<string>

    export function innerHTML(sel: string, ...options: QueryOption[]): ValuedAction<string>

    /**
     * fetch dom node's javascript value
     * @param sel selector
     * @param options options
     */
    export function value(sel: string, ...options: QueryOption[]): ValuedAction<string>

    export function text(sel: string, ...options: QueryOption[]): ValuedAction<string>

    export function textContent(sel: string, ...options: QueryOption[]): ValuedAction<string>

    export type Quad = [number, number, number, number]

    export interface DomBoxMode {
        content: Quad
        padding: Quad
        border: Quad
        margin: Quad
        width: number
        height: number
        shapeOutside?: ShapeOutsideInfo

    }

    export interface ShapeOutsideInfo {
        bounds: Quad
        shape: Uint8Array
        marginShape: Uint8Array
    }

    export function dimensions(sel: string, ...options: QueryOption[]): ValuedAction<DomBoxMode>

    export function attributes(sel: string, ...options: QueryOption[]): ValuedAction<Record<string, string>>

    export function attributesAll(sel: string, ...options: QueryOption[]): ValuedAction<Array<Record<string, string>>>

    export function setValue(sel, value: string, ...options: QueryOption[]): Action

    export function setAttributes(sel: string, attributes: Record<string, string>, ...options: QueryOption[]): Action

    export function setAttributeValue(sel, name, value: string, attributes: Record<string, string>, ...options: QueryOption[]): Action

    export function removeAttribute(sel, name: string, attributes: Record<string, string>, ...options: QueryOption[]): Action

    export function javascriptAttribute(sel, name: string, ...options: QueryOption[]): ValuedAction<any>

    export function setJavascriptAttribute(sel, name: string, value: string, ...options: QueryOption[]): Action

    export function sendKeys(sel, keyCode: string, ...options: QueryOption[]): Action

    export function setUploadFiles(sel: string, files: string[], ...options: QueryOption[]): Action

    export interface ComputedStyle {
        name: string
        value: string
    }

    export function computedStyle(sel: string, ...options: QueryOption[]): ValuedAction<ComputedStyle[]>

    export interface CssSourceRange {
        startLine: number
        startColumn: number
        endLine: number
        endColumn: number
    }

    export interface CssProperty {
        name: string
        value: string
        important?: boolean
        implicit?: boolean
        text?: string
        parsedOk?: boolean
        disabled?: boolean
        range?: CssSourceRange
        longhandProperties?: CssProperty[]
    }

    export interface CssShorthandEntry {
        name: string
        value: string
        important?: boolean
    }

    export interface Style {
        styleSheetID?: string
        cssProperties: CssProperty[]
        shorthandEntries: CssShorthandEntry[]
        cssText?: string
        range?: CssSourceRange
    }

    export interface RuleMatch {
        rule: Rule
        matchingSelectors: number[]
    }

    export interface Rule {
    }

    export interface PseudoElementMatches {

    }

    export interface InheritedStyleEntry {

    }

    export interface InheritedPseudoElementMatches {

    }

    export interface KeyframesRule {

    }

    export interface PositionFallbackRule {

    }

    export interface MatchedStyle {
        inlineStyle?: Style
        attributesStyle?: Style
        matchedCSSRules?: RuleMatch[]
        pseudoElements?: PseudoElementMatches[]
        inherited?: InheritedStyleEntry[]
        inheritedPseudoElements?: InheritedPseudoElementMatches[]
        cSSKeyframesRules?: KeyframesRule[]
        cSSPositionFallbackRules?: PositionFallbackRule[]
        parentLayoutNodeID?: number
    }

    export function matchedStyle(sel: string, ...options: QueryOption[]): ValuedAction<MatchedStyle>

    export function screenShot(sel: string, ...options: QueryOption[]): ValuedAction<Uint8Array>

    export function captureScreenshot(): ValuedAction<Uint8Array>

    /**
     *
     * @param quality [0 .. 100 ]
     */
    export function fullScreenshot(quality: number): ValuedAction<Uint8Array>

    export interface EvaluateOption {
    }

    export const EvalWithCommandLineAPI: EvaluateOption
    export const EvalIgnoreExceptions: EvaluateOption
    export const EvalAsValue: EvaluateOption

    export function evalObjectGroup(objectGroup: string): EvaluateOption

    export function evaluate(expr: string, ...options: EvaluateOption[]): ValuedAction<any>
    export function evaluateAsDevTools(expr: string, ...options: EvaluateOption[]): ValuedAction<any>
    export function callFunctionOn(func: string, ...args:any[]): ValuedAction<any>



}
