import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { HomeScreen } from '../screens/HomeScreen';
import { VoiceRecorderDetailsScreen } from '../screens/voiceRecorder/VoiceRecorderDetailsScreen';

const { Navigator, Screen } = createNativeStackNavigator();

export const MainNavigator = () => (
    <Navigator screenOptions={{ headerShown: false }}>
        <Screen name='Home' component={HomeScreen} />
        <Screen name='Details' component={VoiceRecorderDetailsScreen} />
    </Navigator>
);