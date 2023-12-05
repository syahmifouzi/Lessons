import { StyleSheet } from 'react-native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { BottomNavigation, BottomNavigationTab, Layout, Text, Icon } from '@ui-kitten/components';

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

const OrdersScreen = () => (
    <Layout style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <Text category='h1'>ORDERS</Text>
    </Layout>
);

const BottomTabBar = ({ navigation, state }) => (
    <BottomNavigation
        selectedIndex={state.index}
        onSelect={index => navigation.navigate(state.routeNames[index])}
    >
        <BottomNavigationTab
            title='USERS'
            icon={PersonIcon}
        />
        <BottomNavigationTab
            title='ORDERS'
            icon={BellIcon}
        />
    </BottomNavigation>
);

export const HomeNavigator = () => (
    <Navigator tabBar={props => <BottomTabBar {...props} />} screenOptions={{ headerShown: false }}>
        <Screen name='Users' component={UsersScreen} />
        <Screen name='Orders' component={OrdersScreen} />
    </Navigator>
);