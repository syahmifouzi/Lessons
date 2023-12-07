import React, { useContext } from 'react';
import { Divider, Icon, Layout, Text, TopNavigation, TopNavigationAction } from '@ui-kitten/components';
import { SafeAreaView } from 'react-native-safe-area-context';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite"

export const VoiceRecorderDetailsScreen = ({ navigation }) => {
    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;

    const navigateBack = () => {
        navigation.goBack();
    };

    const BackIcon = (props) => (
        <Icon {...props} name='arrow-back' />
    );
    const EditIcon = (props) => (
        <Icon {...props} name='edit' />
    );
    const DeleteIcon = (props) => (
        <Icon {...props} name='trash-2' />
    );

    const BackAction = () => (
        <TopNavigationAction icon={BackIcon} onPress={navigateBack} />
    );

    const MoreActions = () => (
        <Layout style={{ flexDirection: "row" }}>
            <TopNavigationAction icon={EditIcon} onPress={navigateBack} />
            <TopNavigationAction icon={DeleteIcon} onPress={navigateBack} />
        </Layout>
    );


    const TopNav = observer(() => {
        let item = audioStore.getSelectedItem;
        return (
            <TopNavigation title={item.title} alignment='center' accessoryLeft={BackAction} accessoryRight={MoreActions} />
        )
    });

    return (
        <Layout style={{ flex: 1 }}>
            <SafeAreaView style={{ flex: 1 }}>
                <TopNav />
                <Divider />
                <Layout style={{
                    flex: 1, justifyContent: 'center', alignItems: 'center'
                }}>
                    <Text>HELLO</Text>
                </Layout>
            </SafeAreaView>
        </Layout>
    );
}