import { Button, Divider, Layout, TopNavigation } from '@ui-kitten/components';
import { HomeNavigator } from '../navigations/HomeNavigator';
import {
    useSafeAreaInsets,
} from 'react-native-safe-area-context';

export const HomeScreen = ({ navigation }) => {
    const insets = useSafeAreaInsets();

    const navigateDetails = () => {
        navigation.navigate('Details');
    };

    return (
        <HomeNavigator />
    );
};