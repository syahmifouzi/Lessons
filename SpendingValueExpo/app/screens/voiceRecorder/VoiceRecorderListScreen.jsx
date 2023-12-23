import React, { useContext, useEffect } from 'react';
import { StyleSheet } from 'react-native';
import { Layout, Text, List, ListItem, Divider, Icon, Button } from '@ui-kitten/components';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite"
import { SafeAreaView } from 'react-native-safe-area-context';

export const VoiceRecorderListScreen = ({ navigation }) => {

    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;
    const firebaseApp = mobxCtx.firebaseApp;


    useEffect(() => {
        // console.log('Run once..');
        audioStore.setListFromServer(firebaseApp);
    }, []);


    const renderIcon = () => (
        <Icon
            style={styles.icon}
            fill='#8F9BB3'
            name='arrow-ios-forward'
        />
    );

    const renderDescription = (item) => {
        return (
            <Text>{item.status} | {item.date()}</Text>
        );
    }

    function goToDetailsScreen(index) {
        audioStore.setSelectedIndex(index);
        navigation.navigate('Details');
    }

    const renderItem = (item) => {
        let dataitem = item.item;
        let index = item.index;
        return (
            <ListItem
                title={dataitem.title}
                description={renderDescription(dataitem)}
                accessoryRight={renderIcon}
                onPress={() => goToDetailsScreen(index)}
            />
        )
    };

    const ListRecording = observer(({ store }) =>
        <List
            style={{}}
            keyExtractor={item => item.id}
            data={store.getListItem}
            renderItem={renderItem}
            ItemSeparatorComponent={Divider}
        />);

    const onAddNew = () => {
        navigation.navigate('Create');
    }


    return (
        <Layout style={{ flex: 1 }}>
            <SafeAreaView style={{ flex: 1 }}>
                <Layout style={{ flexDirection: 'row' }}>
                    <Text category='h1' style={{ marginRight: 20 }}>List</Text>
                    <Button onPress={onAddNew} >Add New</Button>
                </Layout>
                <ListRecording store={audioStore} />
            </SafeAreaView>
        </Layout>
    );
}

const styles = StyleSheet.create({
    icon: {
        width: 32,
        height: 32,
    },
});