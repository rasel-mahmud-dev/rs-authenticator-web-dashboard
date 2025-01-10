import {create} from 'zustand';

const useAuthStore = create((set) => ({
    authLoaded: false,
    user: null,

    setAuth: (userData, token) => set({authLoaded: true, user: userData}),
    logout: () => set({authLoaded: true, user: null}),
    updateUser: (newUserData) => set({user: {...newUserData}}),
}));

export default useAuthStore;
