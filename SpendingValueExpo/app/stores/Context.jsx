import { createContext } from 'react';
import { AudioStore } from './AudioStore';
import { firebaseApp } from './FirebaseInit';

class ContextStores {
    audioStore = AudioStore()
    firebaseApp = firebaseApp;
}

const passContext = new ContextStores()
export const MobxContext = createContext(passContext);