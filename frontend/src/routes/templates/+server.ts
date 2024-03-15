import { json } from '@sveltejs/kit';
import { Backend } from '$lib/backend-service/autoreplace';


export async function POST({ request }) {
	const template = await request.json();

	var createdTemplate = await Backend.createTemplate(template)

	return json(createdTemplate, { status: 201 })
}
