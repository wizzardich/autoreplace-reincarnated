import type { PageServerLoad, RequestEvent } from './$types';
import { env } from '$env/dynamic/private';
import { BackendService } from '$lib/backend/autoreplace';

const backend = new BackendService(env.SVC_AUTOREPLACE_BASE_URL);

export const load: PageServerLoad = async ({ }) => {
    const templateList = await backend.getTemplates();
    return { templates: templateList };
};
