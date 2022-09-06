export namespace main {
	
	export class Action {
	    icon: string;
	    title: string;
	    type: string;
	    content: string;
	    path: string;
	    url: string;
	    args: string[];
	
	    static createFrom(source: any = {}) {
	        return new Action(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.icon = source["icon"];
	        this.title = source["title"];
	        this.type = source["type"];
	        this.content = source["content"];
	        this.path = source["path"];
	        this.url = source["url"];
	        this.args = source["args"];
	    }
	}
	export class SearchItem {
	    icon_src: string;
	    title: string;
	    subtitle: string;
	    fill: string;
	    accessory_title: string;
	    keywords: string[];
	    actions: Action[];
	
	    static createFrom(source: any = {}) {
	        return new SearchItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.icon_src = source["icon_src"];
	        this.title = source["title"];
	        this.subtitle = source["subtitle"];
	        this.fill = source["fill"];
	        this.accessory_title = source["accessory_title"];
	        this.keywords = source["keywords"];
	        this.actions = this.convertValues(source["actions"], Action);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

