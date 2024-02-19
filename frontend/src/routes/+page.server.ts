import type { PageServerLoad, RequestEvent } from './$types';
import { env } from '$env/dynamic/private';
import { BackendService } from '$lib/backend/autoreplace';

const backend = new BackendService(env.SVC_AUTOREPLACE_BASE_URL);

export const load: PageServerLoad = async ({ }) => {
    const templateList = await backend.getTemplates();
	console.log(templateList);
	templateList.forEach((template) => {
		console.log(template);
		template.replacements.forEach((replacement) => {
			console.log(replacement);
		});
	});

    return { templates: templateList };
};

// export const actions = {
// 	//@ts-ignore
// 	loadTemplate: async ({ request }) => {
// 		const data = await request.formData();
// 		const template = await backend.fetchTemplateById(data.get('template-id'));
// 		console.log(template);
// 		return { success: true, template: template };
// 	},
// };
