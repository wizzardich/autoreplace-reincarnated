import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
    const templateListResponse = await fetch(`http://localhost:8090/templates`);
    const templateList = await templateListResponse.json();

    return { templates: templateList };
};
