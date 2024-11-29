import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyB3Ye-kOJU1-C-9utCQNVwyi-ULBI4QWe4",
  authDomain: "book-my-hotel-7b670.firebaseapp.com",
  projectId: "book-my-hotel-7b670",
  storageBucket: "book-my-hotel-7b670.appspot.com",
  messagingSenderId: "129728489120",
  appId: "1:129728489120:web:621f6d921365c97a304bc5",
  measurementId: "G-FP2BLPC8FD"
};

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

export { auth };
export default app;
