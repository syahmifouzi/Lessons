import React from 'react';
import { Button, Divider, Layout, TopNavigation } from '@ui-kitten/components';
import {
    useSafeAreaInsets,
} from 'react-native-safe-area-context';

export const HomeScreen = ({ navigation }) => {
    const insets = useSafeAreaInsets();

    const navigateDetails = () => {
        navigation.navigate('Details');
    };

    return (
        <Layout style={{
            flex: 1,
            paddingTop: insets.top,
            paddingBottom: insets.bottom,
            paddingLeft: insets.left,
            paddingRight: insets.right,
        }}>
            <TopNavigation title='MyApp' alignment='center' />
            <Divider />
            <Layout style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
                <Button onPress={navigateDetails}>OPEN DETAILS</Button>
            </Layout>
        </Layout>
    );
};