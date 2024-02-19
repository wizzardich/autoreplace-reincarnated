class Replacement {
    from: string;
    to: string;

    constructor(from: string, to: string) {
        this.from = from;
        this.to = to;
    }
}

class Template {
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