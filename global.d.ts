// @ts-ignore
import io from 'go/io'

/**
 * register a resource to auto closed after execution
 * @param closer
 */
declare function registerResource<T extends io.Closer>(closer:T):T
/**
 * unregister a resource for already closed
 * @param closer
 */
declare function removeResource<T extends io.Closer>(closer:T):T