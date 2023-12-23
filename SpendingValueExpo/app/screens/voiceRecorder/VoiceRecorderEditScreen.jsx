import React, { useContext, useEffect } from 'react';
import { StyleSheet, View, TextInput, ScrollView } from 'react-native';
import { Divider, Icon, Layout, Text, TopNavigation, TopNavigationAction, Modal, Card, Button } from '@ui-kitten/components';
import { SafeAreaView } from 'react-native-safe-area-context';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite"
import moment from 'moment';
import DateTimePicker from '@react-native-community/datetimepicker';

export const VoiceRecorderEditScreen = ({ navigation }) => {
    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;
    const firebaseApp = mobxCtx.firebaseApp;

    const [visible, setVisible] = React.useState(false);

    const [titleInput, setTitleInput] = React.useState('');
    const [statusInput, setStatusInput] = React.useState('');
    const [dateInput, setDateInput] = React.useState('');
    const [surahInput, setSurahInput] = React.useState('');
    const [ayatStartInput, setAyatStartInput] = React.useState('');
    const [ayatEndInput, setAyatEndInput] = React.useState('');

    const [date, setDate] = React.useState(new Date(1598051730000));
    const [mode, setMode] = React.useState('date');
    const [show, setShow] = React.useState(false);

    useEffect(() => {
        // console.log('Run once..');
        let item = audioStore.getSelectedItem;
        let xx = item.timestamp;
        let xx2 = xx.toDate();
        // let xx3 = Date.parse(xx2.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }));
        let momentbe = moment(xx2.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }), "MM/DD/YYYY, hh:mm:ss A");
        setDate(new Date(momentbe.toLocaleString()));
        // let momentbe = moment("1995-12-25");
        // console.log(xx2.toLocaleString("en-US", { timeZone: 'Asia/Kuala_Lumpur' }));
        // console.log(momentbe);
        setTitleInput(item.title);
        setStatusInput(item.status);
        setDateInput(momentbe);
        setSurahInput(item.surah);
        setAyatStartInput(item.ayatStart);
        setAyatEndInput(item.ayatEnd);
    }, []);

    const navigateBack = () => {
        navigation.goBack();
    };

    const BackIcon = (props) => (
        <Icon {...props} name='arrow-back' />
    );
    const EditIcon = (props) => (
        <Text>Save</Text>
    );

    const BackAction = () => (
        <TopNavigationAction icon={BackIcon} onPress={navigateBack} />
    );

    const MoreActions = () => (
        <Layout style={{ flexDirection: "row" }}>
            <TopNavigationAction icon={EditIcon} onPress={async () => {
                let data = {
                    title: titleInput,
                    status: statusInput,
                    surah: surahInput,
                    ayatStart: ayatStartInput,
                    ayatEnd: ayatEndInput,
                }
                await audioStore.updateItemOnDB(firebaseApp, data);
                navigateBack();
            }} />
        </Layout>
    );

    const TopNav = observer(() => {
        let item = audioStore.getSelectedItem;
        return (
            <TopNavigation title={`Edit ${item.title}`} alignment='center' accessoryLeft={BackAction} accessoryRight={MoreActions} />
        )
    });

    const showMode = (currentMode) => {
        setShow(true);
        setMode(currentMode);
    };

    const onChange = (event, selectedDate) => {
        const currentDate = selectedDate;
        setShow(false);
        setDate(currentDate);
    };

    const showDatepicker = () => {
        showMode('date');
    };

    const showTimepicker = () => {
        showMode('time');
    };

    return (
        <Layout style={{ flex: 1 }}>
            <ScrollView>
                <SafeAreaView style={{ flex: 1 }}>
                    <TopNav />
                    <Divider />
                    <Layout style={{
                        flex: 1, justifyContent: 'center', alignItems: 'center'
                    }}>
                        <Layout style={{ padding: 10, width: '100%' }}>
                            <TextInput
                                style={styles.input}
                                onChangeText={setTitleInput}
                                value={titleInput}
                                placeholder="Title"
                            />
                        </Layout>
                        <Layout style={{ padding: 10, width: '100%' }}>
                            <TextInput
                                style={styles.input}
                                onChangeText={setStatusInput}
                                value={statusInput}
                                placeholder="Status"
                            />
                        </Layout>
                        {/* <Text>selected: {date.toLocaleString()}</Text>
                        <Button onPress={showDatepicker} title="Show date picker!" />
                        {show && (
                            <DateTimePicker
                                testID="dateTimePicker"
                                value={date}
                                mode={mode}
                                is24Hour={true}
                                onChange={onChange}
                            />
                        )} */}
                        <Layout style={{ padding: 10, width: '100%' }}>
                            <TextInput
                                style={styles.input}
                                onChangeText={setSurahInput}
                                value={surahInput}
                                placeholder="Surah"
                            />
                        </Layout>
                        <Layout style={{ padding: 10, width: '100%' }}>
                            <TextInput
                                style={styles.input}
                                onChangeText={setAyatStartInput}
                                value={ayatStartInput}
                                placeholder="Ayat Start"
                            />
                        </Layout>
                        <Layout style={{ padding: 10, width: '100%' }}>
                            <TextInput
                                style={styles.input}
                                onChangeText={setAyatEndInput}
                                value={ayatEndInput}
                                placeholder="Ayat End"
                            />
                        </Layout>

                        <Modal
                            visible={visible}
                            backdropStyle={styles.backdrop}
                            onBackdropPress={() => setVisible(false)}
                        >
                            <Card disabled={true}>
                                <Text>Hello</Text>
                            </Card>
                        </Modal>
                    </Layout>
                </SafeAreaView>
            </ScrollView>
        </Layout>
    );
}
const styles = StyleSheet.create({
    container: {
        minHeight: 192,
    },
    backdrop: {
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
    },
    input: {
        height: 40,
        borderWidth: 1,
        padding: 10,
    },
});