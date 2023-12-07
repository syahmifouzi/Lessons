import { initializeApp } from 'firebase/app';

const firebaseConfig = {
    apiKey: "AIzaSyB903AfKNnJG1JpQVKj9pCfUZMZLdQaaS4",
    authDomain: "spending-value.firebaseapp.com",
    projectId: "spending-value",
    storageBucket: "spending-value.appspot.com",
    messagingSenderId: "972509417548",
    appId: "1:972509417548:web:8f77d6b674ff2356b2a4e9",
    measurementId: "G-RC5R35CXMJ"
};

const firebaseApp = initializeApp(firebaseConfig);

export { firebaseApp };