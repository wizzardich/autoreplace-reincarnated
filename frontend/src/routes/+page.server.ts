import type { PageServerLoad, RequestEvent } from './$types';
import { Backend } from '$lib/backend/autoreplace';

export const load: PageServerLoad = async ({ }) => {
    const templateList = await Backend.getTemplates();
    return { templates: templateList };
};
