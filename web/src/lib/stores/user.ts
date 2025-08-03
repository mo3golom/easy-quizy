import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { User } from '../types';


const defaultUser: User = {
    id: Math.floor(Math.random() * Number.MAX_SAFE_INTEGER),
    username: "default",
    source: "browser",
};

export const userData = writable<User>(defaultUser);

export function restoreUser() {
    if (browser) {
        const storedUser = localStorage.getItem('userData');
        if (storedUser) {
            const user: User = JSON.parse(storedUser);
            userData.set(user);
        } else {
            localStorage.setItem('userData', JSON.stringify(defaultUser));
        }
    }
}

export function setTelegramUser(user: any, chatId?: number, chatType?: string) {
    const telegramUser: User = {
        id: user.id,
        username: user.username,
        language_code: user.language_code,
        source: "telegram",

    }

    if (chatId) {
        telegramUser.chatId = chatId;
    }
    if (chatType) {
        telegramUser.chatType = chatType;
    }
    
    userData.set(telegramUser);

    if (browser) {
        localStorage.setItem('userData', JSON.stringify(telegramUser));
    }
}

// Функция для отладки (можно вызвать в консоли браузера)
export function debugUser() {
    console.log('Current user in store:', userData);
    if (browser) {
        console.log('User in localStorage:', localStorage.getItem('userData'));
    }
}