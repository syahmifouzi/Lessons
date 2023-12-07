import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { BottomNavigation, BottomNavigationTab, Layout, Text, Icon } from '@ui-kitten/components';
import { VoiceRecorderScreen } from '../screens/voiceRecorder/VoiceRecorderScreen';
import { VoiceRecorderListScreen } from '../screens/voiceRecorder/VoiceRecorderListScreen';

const { Navigator, Screen } = createBottomTabNavigator();

const PersonIcon = (props) => (
    <Icon
        {...props}
        name='person-outline'
    />
);

const BellIcon = (props) => (
    <Icon
        {...props}
        name='bell-outline'
    />
);

const UsersScreen = () => (
    <Layout style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text category='h1'>USERS</Text>
    </Layout>
);

const BottomTabBar = ({ navigation, state }) => (
    <BottomNavigation
        selectedIndex={state.index}
        onSelect={index => navigation.navigate(state.routeNames[index])}
    >
        <BottomNavigationTab
            title='Users'
            icon={PersonIcon}
        />
        <BottomNavigationTab
            title='Orders'
            icon={BellIcon}
        />
        <BottomNavigationTab
            title='VoiceList'
            icon={BellIcon}
        />
    </BottomNavigation>
);

export const HomeNavigator = () => (
    <Navigator tabBar={props => <BottomTabBar {...props} />} screenOptions={{ headerShown: false }}>
        <Screen name='Users' component={UsersScreen} />
        <Screen name='Orders' component={VoiceRecorderScreen} />
        <Screen name='VoiceList' component={VoiceRecorderListScreen} />
    </Navigator>
);