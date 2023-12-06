import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { HomeScreen } from '../screens/HomeScreen';
// import { DetailsScreen } from '../testkitty/details';

const { Navigator, Screen } = createNativeStackNavigator();

export const MainNavigator = () => (
    <Navigator screenOptions={{ headerShown: false }}>
        <Screen name='Home' component={HomeScreen} />
        {/* <Screen name='Details' component={DetailsScreen} /> */}
    </Navigator>
);