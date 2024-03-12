import { json } from '@sveltejs/kit';
import { Backend } from '$lib/backend-service/autoreplace';


export async function POST({ request }) {
	const template = await request.json();

	await Backend.createTemplate(template)

	return new Response(null, { status: 201 })
}
