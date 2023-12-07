import React, { useState, useEffect, useContext } from 'react';
import { ScrollView } from 'react-native';
import { Layout, Text, Button } from '@ui-kitten/components';
import { Alert } from 'react-native';
import { Audio } from 'expo-av';
import * as FileSystem from 'expo-file-system';
import { Slider } from '@rneui/themed';
import { MobxContext } from '../../stores/Context';
import { observer } from "mobx-react-lite";
import { SafeAreaView } from 'react-native-safe-area-context';

export const VoiceRecorderScreen = () => {
    const [recording, setRecording] = useState();
    const [sound, setSound] = useState();
    const [playbackuri, setPlaybackuri] = useState();
    const [slidervalue, setSlidervalue] = useState({
        value: 0,
        duration: 5000
    });

    const mobxCtx = useContext(MobxContext);
    const audioStore = mobxCtx.audioStore;

    useEffect(() => {
        console.log('Run once..');
    }, []);

    const _onRecordingStatusUpdate = (recordingStatus) => {
        // console.log(`Recording update`);
        // console.log(recordingStatus);
    }

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

        // Create a file name for the recording
        const fileName = `recording-${Date.now()}.m4a`;
        const newDir = `${FileSystem.documentDirectory}recordings/${fileName}`;

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
        console.log(`New URI: ${newDir}`);
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
            await sound.loadAsync(
                { uri: playbackuri },
                // { uri: "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3" },
                // { shouldPlay: true, }
            );
            // const { playbackObject } = await Audio.Sound.createAsync(
            //     { uri: "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3" },
            //     { shouldPlay: true,}
            //     );
            // await sound.setPositionAsync(0);
            sound.setOnPlaybackStatusUpdate(_onPlaybackStatusUpdate);
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

    // React.useEffect(() => {
    //     return sound
    //         ? () => {
    //             console.log('Unloading Sound');
    //             sound.unloadAsync();
    //         }
    //         : undefined;
    // }, [sound]);

    const ListLength = observer(({ store }) => <Text>Length is: {store.getListItem.length}</Text>);

    return (

        <Layout style={{ flex: 1 }}>
            <ScrollView>
                <SafeAreaView style={{ flex: 1 }}>
                    <Layout style={{
                        flex: 1, justifyContent: 'center', alignItems: 'center'
                    }}>
                        <Text category='h1'>Voice Recorder</Text>
                        <ListLength store={audioStore} />
                        <Button onPress={() => Alert.alert('Alert Title')} >Alert</Button>
                        <Button onPress={startRecording} >Start Recording</Button>
                        <Text category='h1'>.</Text>
                        <Button onPress={stopRecording} >Stop Recording</Button>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Text category='h1'>.</Text>
                        <Button onPress={playSound} >Playback Recording</Button>
                        <Text category='h1'>.</Text>
                        <Button onPress={stopSound} >Stop Playback Recording</Button>
                        <Text category='h1'>.</Text>
                        <Layout style={{ padding: 20, justifyContent: 'center', alignItems: 'stretch', width: '100%' }}>
                            <Text>{_millisToDuration(slidervalue.value)}</Text>
                            <Slider
                                value={slidervalue.value}
                                onValueChange={_onSliderValueChanged}
                                minimumValue={0}
                                maximumValue={slidervalue.duration}
                                step={1}
                            />
                            <Text>{_millisToDuration(slidervalue.duration)}</Text>
                        </Layout>
                    </Layout>
                </SafeAreaView>
            </ScrollView>
        </Layout>
    );
}