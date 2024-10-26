import { fail, redirect } from '@sveltejs/kit';
import { userStore } from '$lib/store';

export const actions = {
    default: async (event) => {
        const formData = await event.request.formData();
        const code = formData.get("code") as string | null;
        const name = formData.get("name") as string | null;

        if (!code) {
            return fail(400, { message: "code is missing" });
        }
        if (!name) {
            return fail(400, { message: "name is missing" });
        }

        // TODO: connecting to the lobby/server logic

        userStore.set({ code, name });

        throw redirect(302, '/lobby');
    },
};