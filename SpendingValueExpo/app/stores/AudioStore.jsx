import { action, makeAutoObservable } from "mobx"
import { getFirestore, collection, getDocs } from "firebase/firestore";

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

    setListItem = action((val)=>{
        this.listItem = val;
    });

    setSelectedIndex = action((val)=>{
        this.selectedIndex = val;
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
        const querySnapshot = await getDocs(collection(db, "audio"));
        return querySnapshot;
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
    datetime = "";

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
            this.datetime = data["datetime"]
        } catch (error) {
            console.error(`Err converting data from DB: ${error}`);
        }
    }

}


export function AudioStore() { return new AudioStoreTable() };