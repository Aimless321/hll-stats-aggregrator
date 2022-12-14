import { defineStore } from 'pinia'

export const useStore = defineStore('main', {
    state: () => ({ loggedIn: localStorage.getItem('username')}),
    actions: {
        login() {
            this.loggedIn = true;
        },
    },
})
