import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { HomeScreen } from '../screens/HomeScreen';
import { VoiceRecorderDetailsScreen } from '../screens/voiceRecorder/VoiceRecorderDetailsScreen';
import { VoiceRecorderEditScreen } from '../screens/voiceRecorder/VoiceRecorderEditScreen';
import { VoiceRecorderCreateScreen } from '../screens/voiceRecorder/VoiceRecorderCreateScreen';

const { Navigator, Screen } = createNativeStackNavigator();

export const MainNavigator = () => (
    <Navigator screenOptions={{ headerShown: false }}>
        <Screen name='Home' component={HomeScreen} />
        <Screen name='Details' component={VoiceRecorderDetailsScreen} />
        <Screen name='Edit' component={VoiceRecorderEditScreen} />
        <Screen name='Create' component={VoiceRecorderCreateScreen} />
    </Navigator>
);