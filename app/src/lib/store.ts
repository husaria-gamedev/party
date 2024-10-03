import { writable } from 'svelte/store';
type UserData = {
    code: string;
    name: string;
};
export const userStore = writable<UserData>({code: '', name:''});
