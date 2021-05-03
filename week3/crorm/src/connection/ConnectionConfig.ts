export interface ConnectionConfig {
    readonly driver : string;
    readonly host : string;
    readonly post : bigint;
    readonly username : string;
    readonly userpassword : string;
    readonly tableName : string;
    readonly database : string;
}

export class Connection {
    readonly driver : string;
    readonly host : string;
    readonly post : bigint;
    readonly username : string;
    readonly userpassword : string;
    readonly tableName : string;
    readonly database : string;
    constructor(options : ConnectionConfig) {
        this.driver = options.driver
        this.host = options.host
        this.post = options.post
        this.username = options.username
        this.userpassword = options.userpassword
        this.tableName = options.tableName
        this.database = options.database;
    }
}