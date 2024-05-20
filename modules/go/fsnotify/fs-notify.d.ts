declare module 'go/fsNotify' {
    export class Watcher {
        //buf: the event buffer size
        constructor(buf?: number)


        // close the Notifier
        close()

        //add path to watch
        add(target: string)

        //add path tree
        addTree(target: string)

        //remove path
        remove(target: string)

        //remove a path tree
        removeTree(target: string)

        //path those are under watch. undefined if already closed
        watchList(): string[] | undefined

        //register callback for events
        onEvent(fn: (ev: FsEvent, er: Error) => void)

        //block and await all events, unless VM is shutdown
        await()
    }

    export interface FsEvent {
        //relative path for a registered target
        path(): string

        //returns the target is a file or directory (if target not exists or permission not granted, also returns false)
        isFile(): boolean

        // target is created
        isCreated(): boolean

        //target is content modified
        isModified(): boolean

        //target is renamed
        isRenamed(): boolean

        //target is removed
        isRemoved(): boolean

        //target's attributes changed
        isAttributeChanged(): boolean
    }
}