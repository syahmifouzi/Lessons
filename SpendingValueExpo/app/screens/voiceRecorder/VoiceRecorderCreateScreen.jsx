import React, { useContext, useEffect } from 'react';
import { StyleSheet, View, TextInput, ScrollView } from 'react-native';
import { Divider, Icon, Layout, Text, TopNavigation, TopNavigationAction, Modal, Card, Button } from '@ui-kitten/components';
import { Audio } from 'expo-av';
import { SafeAreaView } from 'react-native-safe-area-context';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite"
import moment from 'moment';

export const VoiceRecorderCreateScreen = ({ navigation }) => {
    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;
    const firebaseApp = mobxCtx.firebaseApp;

    const [visible, setVisible] = React.useState(false);

    const [titleInput, setTitleInput] = React.useState('');
    const [surahInput, setSurahInput] = React.useState('');
    const [ayatStartInput, setAyatStartInput] = React.useState('');
    const [ayatEndInput, setAyatEndInput] = React.useState('');

    const [recording, setRecording] = React.useState();
    const [playbackuri, setPlaybackuri] = React.useState();
    const [recordingText, setRecordingText] = React.useState('Idle');

    useEffect(() => {
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

    const onSaveButton = async () => {
        let data = {
            title: titleInput,
            status: 'Pending',
            surah: surahInput,
            ayatStart: ayatStartInput,
            ayatEnd: ayatEndInput,
            timestamp: new Date(Date.now()),
            duration: 'TBD',
            url: 'TBD',
        }
        await audioStore.createItemOnDB(firebaseApp, data, playbackuri);
        navigateBack();

    }

    const MoreActions = () => (
        <Layout style={{ flexDirection: "row" }}>
            <TopNavigationAction icon={EditIcon} onPress={onSaveButton} />
        </Layout>
    );

    const TopNav = observer(() => {
        let item = audioStore.getSelectedItem;
        return (
            <TopNavigation title={'Create New'} alignment='center' accessoryLeft={BackAction} accessoryRight={MoreActions} />
        )
    });

    const _onRecordingStatusUpdate = (recordingStatus) => {
        // console.log(`Recording update`);
        // console.log(recordingStatus);
    }

    async function startRecording() {
        try {
            console.log('Requesting permissions..');
            await Audio.requestPermissionsAsync();
            await Audio.setAudioModeAsync({
                allowsRecordingIOS: true,
                playsInSilentModeIOS: true,
            });

            console.log('Starting recording..');
            const { recording } = await Audio.Recording.createAsync(Audio.RecordingOptionsPresets.HIGH_QUALITY,
                _onRecordingStatusUpdate
            );
            setRecording(recording);
            console.log('Recording started');
            setRecordingText('Now Recording...');
            // Which is equivalent to the following:
            // const recording = new Audio.Recording();
            // await recording.prepareToRecordAsync(options);
            // recording.setOnRecordingStatusUpdate(onRecordingStatusUpdate);
            // await recording.startAsync();
        } catch (err) {
            console.error('Failed to start recording', err);
        }
    }

    async function stopRecording() {
        console.log('Stopping recording..');
        setRecording(undefined);
        await recording.stopAndUnloadAsync();
        await Audio.setAudioModeAsync(
            {
                allowsRecordingIOS: false,
            }
        );
        const uri = recording.getURI();
        console.log('Recording stopped and stored at', uri);
        setRecordingText('Recording Stopped');

        // Create a file name for the recording
        // const fileName = `recording-${Date.now()}.m4a`;
        // const newDir = `${FileSystem.documentDirectory}recordings/${fileName}`;

        // Move the recording to the new directory with the new file name
        // await FileSystem.makeDirectoryAsync(FileSystem.documentDirectory + 'recordings/', { intermediates: true });
        // await FileSystem.moveAsync({
        //   from: uri,
        //   to: newDir
        // });
        // This is for simply playing the sound back
        // const playbackObject = new Audio.Sound();
        // await playbackObject.loadAsync({ uri: uri });
        // await playbackObject.playAsync();
        // await playbackObject.unloadAsync();

        setPlaybackuri(uri);
        // console.log(`New URI: ${FileSystem.documentDirectory}recordings/${fileName}`);
        // console.log(`New URI: ${newDir}`);
    }

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

                        <Text>{recordingText}</Text>


                        <Button onPress={startRecording} >Start Recording</Button>
                        <Text category='h1'>.</Text>
                        <Button onPress={stopRecording} >Stop Recording</Button>

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