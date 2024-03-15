import { env } from '$env/dynamic/private';

import { Template } from '$lib/backend/domain';

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
        fetch(`${this.base}/templates/${template.id}`, {
            method: 'PUT',
            body: JSON.stringify(template)
        }).then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
        })
    }

    async deleteTemplate(id: string) {
        fetch(`${this.base}/templates/${id}`, {
            method: 'DELETE'
        }).then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
        })
    }

    async createTemplate(template: Template) {
        return fetch(`${this.base}/templates/`, {
            method: 'POST',
            body: JSON.stringify(template)
        }).then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
            return response.json() as Promise<Template>
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
