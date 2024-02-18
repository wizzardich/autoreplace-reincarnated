import type { PageLoad } from './$types';

export const load: PageLoad = async ({ data }) => {

    console.log('data', data.templates);

	return {
		templates: data.templates,
	};
};