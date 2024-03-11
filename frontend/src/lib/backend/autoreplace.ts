import { env } from '$env/dynamic/private';

export class Replacement {
    from: string;
    to: string;

    constructor(from: string, to: string) {
        this.from = from;
        this.to = to;
    }
}

export class Template {
    id: string;
    name: string;
    replacements: Replacement[];

    constructor(id: string, name: string, replacements: Replacement[]) {
        this.id = id;
        this.name = name;
        this.replacements = replacements;
    }
}

export class BackendService {
    base: string;

    constructor(base: string) {
        this.base = base;
    }

    async getTemplates() {
        return api<Template[]>(`${this.base}/templates`)
    }

    async fetchTemplateById(id: string) {
        return api<Template>(`${this.base}/templates/${id}`)
    }

    async editTemplate(template: Template) {
        console.log(template)
        fetch(`${this.base}/templates/${template.id}`, {
            method: 'PUT',
            body: JSON.stringify(template)
        }).then(response => {
            console.log(response)
            if (!response.ok) {
                throw new Error(response.statusText)
            }
        })
    }
}

// For the "unwrapping" variation

function api<T>(url: string): Promise<T> {
    return fetch(url)
      .then(response => {
        if (!response.ok) {
          throw new Error(response.statusText)
        }
        return response.json() as Promise<T>
      })
}

export const Backend = new BackendService(env.SVC_AUTOREPLACE_BASE_URL);
