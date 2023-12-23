import { action, makeAutoObservable } from "mobx"
import { getFirestore, collection, doc, addDoc, getDocs, updateDoc, setDoc, deleteDoc, query, orderBy, limit } from "firebase/firestore";
import { getStorage, ref, uploadBytes, getDownloadURL, deleteObject } from "firebase/storage";
import moment from 'moment';

class AudioStoreTable {
    details = new AudioDoc();
    listItem = [];
    selectedIndex = 0;

    constructor() {
        makeAutoObservable(this);
    }

    get getListItem() {
        return this.listItem;
    }

    get getSelectedItem() {
        return this.listItem[this.selectedIndex];
    }

    setListItem = action((val) => {
        this.listItem = val;
    });

    setSelectedIndex = action((val) => {
        this.selectedIndex = val;
    });

    updateSelectedItem = action((val) => {
        let item = this.getSelectedItem;
        item = {
            ...item,
            ...val
        }
    });


    setListFromServer = async (firebaseApp) => {
        const querySnapshot = await this.getListFromDB(firebaseApp);

        let tempList = []
        querySnapshot.forEach((doc) => {
            let audioDoc = new AudioDoc();
            audioDoc.fromDB(doc.id, doc.data());
            tempList.push(audioDoc);
        });
        this.setListItem(tempList);
    }

    getListFromDB = async (firebaseApp) => {
        const db = getFirestore(firebaseApp);
        const audioRef = collection(db, "audio"); 
        const q = query(audioRef, orderBy("timestamp", "desc"), limit(10));
        // const querySnapshot = await getDocs(collection(db, "audio"));
        const querySnapshot = await getDocs(q);
        return querySnapshot;
    }

    updateItemOnDB = async (firebaseApp, data) => {
        let item = this.getSelectedItem;
        const db = getFirestore(firebaseApp);
        const dbref = doc(db, "audio", item.id);
        await updateDoc(dbref, data);
        this.updateSelectedItem(data);
        // let testi = {
        //     one:'one',
        //     two:'two',
        //     three:'three',
        // }
        // let test2 = {
        //     three:'four',
        // }
        // testi = {
        //     ...testi,
        //     ...test2
        // }
        // console.log(testi);

        // refresh listing
        this.setListFromServer(firebaseApp);
    }

    createAudioOnDB = async (firebaseApp, id, file) => {
        const blob = await new Promise((resolve, reject) => {
            const fetchXHR = new XMLHttpRequest();
            fetchXHR.onload = function () {
                resolve(fetchXHR.response);
            };
            fetchXHR.onerror = function (e) {
                reject(new TypeError('Network request failed'));
            };
            fetchXHR.responseType = 'blob';
            fetchXHR.open('GET', file, true);
            fetchXHR.send(null);
        }).catch((err) => console.log(err));
        const storage = getStorage(firebaseApp);
        const storageRef = ref(storage, `audio/${id}.m4a`);

        // uploadBytes(storageRef, file).then((snapshot) => {
        //     console.log('Uploaded a blob or file!');
        //     // console.log(snapshot.ref);
        //     getDownloadURL(snapshot.ref).then((downloadUrl) => {
        //         console.log('File available at', downloadUrl);
        //     })
        // });

        const snapshot = await uploadBytes(storageRef, blob);
        const downloadUrl = await getDownloadURL(snapshot.ref);
        return downloadUrl;

    }

    createItemOnDB = async (firebaseApp, data, file) => {
        const db = getFirestore(firebaseApp);
        // const docRef = await addDoc(collection(db, "audio"), data);
        const docRef = doc(collection(db, "audio"));

        const downloadUrl = await this.createAudioOnDB(firebaseApp, docRef.id, file);

        data.url = downloadUrl;

        await setDoc(docRef, data);

        // refresh listing
        this.setListFromServer(firebaseApp);
    }

    deleteAudioOnDB = async (firebaseApp, id) => {
        const storage = getStorage(firebaseApp);
        const storageRef = ref(storage, `audio/${id}.m4a`);

        deleteObject(storageRef);
    }

    deleteItemOnDB = async (firebaseApp, id) => {
        const db = getFirestore(firebaseApp);
        this.deleteAudioOnDB(firebaseApp, id);
        await deleteDoc(doc(db, "audio", id));

        // refresh listing
        this.setListFromServer(firebaseApp);
    }
}

class AudioDoc {

    id = -1;
    title = "";
    url = "";
    duration = "";
    surah = "";
    status = "";
    ayatStart = "";
    ayatEnd = "";
    timestamp = "";

    fromDB(id, data) {
        this.id = id
        try {
            this.title = data["title"]
            this.url = data["url"]
            this.duration = data["duration"]
            this.surah = data["surah"]
            this.status = data["status"]
            this.ayatStart = data["ayatStart"]
            this.ayatEnd = data["ayatEnd"]
            this.timestamp = data["timestamp"]
        } catch (error) {
            console.error(`Err converting data from DB: ${error}`);
        }
    }

    date() {
        let mydate = this.timestamp.toDate();
        mydate = moment(mydate.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }), "MM/DD/YYYY, hh:mm:ss A");
        mydate = mydate.format("DD/MM/YYYY, HH:mm");

        return mydate;

        // console.log('Run once..');
        // let item = audioStore.getSelectedItem;
        // let xx = item.timestamp;
        // let xx2 = xx.toDate();
        // // let xx3 = Date.parse(xx2.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }));
        // let momentbe = moment(xx2.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }), "MM/DD/YYYY, hh:mm:ss A");
        // // setDate(new Date(momentbe.toLocaleString()));
        // // let momentbe = moment("1995-12-25");
        // console.log(xx2.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }));
        // console.log(momentbe);
    }

}


export function AudioStore() { return new AudioStoreTable() };