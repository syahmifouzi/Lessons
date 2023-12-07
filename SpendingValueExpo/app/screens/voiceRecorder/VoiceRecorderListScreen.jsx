import React, { useContext, useEffect } from 'react';
import { Layout, Text, Button } from '@ui-kitten/components';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite"

export const VoiceRecorderListScreen = () => {

    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;
    const firebaseApp = mobxCtx.firebaseApp;


    useEffect(() => {
        // console.log('Run once..');
        // getListFromDB();
        audioStore.setListFromServer(firebaseApp);
        // console.log(`mobx test: ${mobxCtx.audioStore.details.testTitle}`);
    }, []);

    const ListLength = observer(({store})=><Text>{store.getListItem.length}</Text>);


    return (
        <Layout style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
            <Text category='h1'>List</Text>
            <ListLength store={audioStore} />
        </Layout>
    );
}