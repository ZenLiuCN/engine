declare interface GoError {
    string(): string

    is(error: GoError)
}

type Err = GoError | null
