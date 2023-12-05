import React from 'react';
import { Divider, Icon, Layout, Text, TopNavigation, TopNavigationAction } from '@ui-kitten/components';
import {
    useSafeAreaInsets,
} from 'react-native-safe-area-context';

const BackIcon = (props) => (
    <Icon {...props} name='arrow-back' />
);

export const DetailsScreen = ({ navigation }) => {
    const insets = useSafeAreaInsets();

    const navigateBack = () => {
        navigation.goBack();
    };

    const BackAction = () => (
        <TopNavigationAction icon={BackIcon} onPress={navigateBack} />
    );

    return (
        <Layout style={{
            flex: 1,
            paddingTop: insets.top,
            paddingBottom: insets.bottom,
            paddingLeft: insets.left,
            paddingRight: insets.right,
        }}>
            <TopNavigation title='MyApp' alignment='center' accessoryLeft={BackAction} />
            <Divider />
            <Layout style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
                <Text category='h1'>DETAILS</Text>
            </Layout>
        </Layout>
    );
};