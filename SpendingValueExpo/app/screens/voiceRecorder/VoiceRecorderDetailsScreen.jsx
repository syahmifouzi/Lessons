import React, { useState, useContext } from 'react';
import { StyleSheet, View } from 'react-native';
import { Divider, Icon, Layout, Text, TopNavigation, TopNavigationAction, Modal, Card, Button } from '@ui-kitten/components';
import { SafeAreaView } from 'react-native-safe-area-context';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite"
import Slider from '@react-native-community/slider';
import { Audio } from 'expo-av';

export const VoiceRecorderDetailsScreen = ({ navigation }) => {
    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;
    const firebaseApp = mobxCtx.firebaseApp;

    const [visible, setVisible] = React.useState(false);

    const [sound, setSound] = useState();
    const [slidervalue, setSlidervalue] = useState({
        value: 0,
        duration: 0
    });

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
            <TopNavigationAction icon={EditIcon} onPress={() => {
                navigation.navigate('Edit');
            }} />
            <TopNavigationAction icon={DeleteIcon} onPress={() => setVisible(true)} />
        </Layout>
    );


    const TopNav = observer(() => {
        let item = audioStore.getSelectedItem;
        return (
            <TopNavigation title={item.title} alignment='center' accessoryLeft={BackAction} accessoryRight={MoreActions} />
        )
    });

    const _onPlaybackStatusUpdate = (playbackStatus) => {
        if (!playbackStatus.isLoaded) {
            // not yet loaded
            console.warn(playbackStatus);
            return;
        }

        // Example:
        // {"androidImplementation": "SimpleExoPlayer", "audioPan": 0, "didJustFinish": false, "durationMillis": 425952, "isBuffering": false, "isLoaded": true, "isLooping": false, "isMuted": false, "isPlaying": true, "playableDurationMillis": 87379, "positionMillis": 7660, "progressUpdateIntervalMillis": 500, "rate": 1, "shouldCorrectPitch": false, "shouldPlay": true, "uri": "/examples/mp3/SoundHelix-Song-2.mp3", "volume": 1}

        setSlidervalue({
            ...slidervalue,
            duration: playbackStatus.durationMillis,
            value: playbackStatus.positionMillis
        });

        // console.log('Playback status updated');
        // console.log(playbackStatus);
    }


    const _onSliderValueChanged = async (value) => {
        // setSlidervalue({
        //     ...slidervalue,
        //     value: value
        // });

        await sound.setPositionAsync(value);
    }

    const _millisToDuration = (value) => {
        return new Date(value).toISOString().slice(11, 19);
    }

    async function playSound() {
        // console.log('Loading Sound');
        // const { sound } = await Audio.Sound.createAsync(require('file:///data/user/0/host.exp.exponent/cache/ExperienceData/%2540syahmifouzi%252Fspending-value/Audio/recording-dddcb369-c6ea-456d-b48d-7706a39b2fe3.m4a')
        // );
        // setSound(sound);

        // console.log('Playing Sound');
        // await sound.playAsync();

        // // Don't forget to unload the sound from memory
        // // when you are done using the Sound object
        // await sound.unloadAsync();
        try {
            console.log('Playing recording');
            const sound = new Audio.Sound();
            setSound(sound);
            let item = audioStore.getSelectedItem;
            await sound.loadAsync(
                { uri: item.url },
                // { uri: "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3" },
                // { shouldPlay: true, }
            );
            // const { playbackObject } = await Audio.Sound.createAsync(
            //     { uri: "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3" },
            //     { shouldPlay: true,}
            //     );
            // await sound.setPositionAsync(0);
            sound.setOnPlaybackStatusUpdate(_onPlaybackStatusUpdate);

            await sound.setAudioModeAsync({
                allowsRecordingIOS: false,
                playsInSilentModeIOS: true,
                interruptionModeIOS: Audio.INTERRUPTION_MODE_IOS_DO_NOT_MIX,
                shouldDuckAndroid: true,
                interruptionModeAndroid: Audio.INTERRUPTION_MODE_ANDROID_DO_NOT_MIX,
                playThroughEarpieceAndroid: false,
            });
            await sound.playAsync();
            // await playbackObject.unloadAsync();
        } catch (error) {
            console.error('Failed to play recording', error);
        }
        console.log('Exiting function');

    }

    async function stopSound() {
        try {
            // Don't forget to unload the sound from memory
            // when you are done using the Sound object
            await sound.unloadAsync();
            console.log('Stopped recording');
        } catch (error) {
            console.error('Failed to stop recording', error);
        }
    }

    const onDelete = () => {
        let item = audioStore.getSelectedItem;
        audioStore.deleteItemOnDB(firebaseApp, item.id);
        setVisible(false);
        navigateBack();
    }


    return (
        <Layout style={{ flex: 1 }}>
            <SafeAreaView style={{ flex: 1 }}>
                <TopNav />
                <Divider />
                <Layout style={{
                    flex: 1, justifyContent: 'center', alignItems: 'center'
                }}>
                    <Button onPress={playSound} >Playback Recording</Button>
                    <Text category='h1'>.</Text>
                    <Button onPress={stopSound} >Stop Playback Recording</Button>
                    <Text category='h1'>.</Text>
                    <Layout style={{ padding: 20, justifyContent: 'center', alignItems: 'stretch', width: '100%' }}>
                        <Text>{_millisToDuration(slidervalue.value)}</Text>
                        <Slider
                            // style={{ width: 200, height: 40 }}
                            value={slidervalue.value}
                            onValueChange={_onSliderValueChanged}
                            minimumValue={0}
                            maximumValue={slidervalue.duration}
                            step={1}
                        />
                        <Text>{_millisToDuration(slidervalue.duration)}</Text>
                    </Layout>

                    <Modal
                        visible={visible}
                        backdropStyle={styles.backdrop}
                        onBackdropPress={() => setVisible(false)}
                    >
                        <Card disabled={true}>
                            <Text>Delete?</Text>
                            <Button onPress={() => setVisible(false)} >No</Button>
                            <Layout style={{ height: 20 }} />
                            <Button onPress={onDelete} >Yes</Button>
                        </Card>
                    </Modal>
                </Layout>
            </SafeAreaView>
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
});