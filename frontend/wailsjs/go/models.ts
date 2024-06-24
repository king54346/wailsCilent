export namespace utils {
	
	export class Game {
	    id: number;
	    name: string;
	    size: number;
	    tags: string[];
	    filepath: string;
	
	    static createFrom(source: any = {}) {
	        return new Game(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.size = source["size"];
	        this.tags = source["tags"];
	        this.filepath = source["filepath"];
	    }
	}
	export class Msg {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Msg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}
	export class WorkOrderGameList {
	    id: number;
	    name: string;
	    size: number;
	    status: string;
	    error_msg: string;
	
	    static createFrom(source: any = {}) {
	        return new WorkOrderGameList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.size = source["size"];
	        this.status = source["status"];
	        this.error_msg = source["error_msg"];
	    }
	}
	export class WorkOrderQuery {
	    pageNo: number;
	    pageSize: number;
	    search: string;
	    status: number;
	
	    static createFrom(source: any = {}) {
	        return new WorkOrderQuery(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pageNo = source["pageNo"];
	        this.pageSize = source["pageSize"];
	        this.search = source["search"];
	        this.status = source["status"];
	    }
	}

}

