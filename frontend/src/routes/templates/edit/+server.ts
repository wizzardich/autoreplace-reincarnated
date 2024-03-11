import { json } from '@sveltejs/kit';
import { Backend } from '$lib/backend-service/autoreplace';

export async function PUT({ request }) {
	const template = await request.json();

	await Backend.editTemplate(template)

	return new Response(null, { status: 204 })
}
