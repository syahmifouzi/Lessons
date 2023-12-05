import { Link, Stack, router } from 'expo-router';
import { Text, View } from 'react-native';

export default function Hello2Page() {
    return (
        <View>
            <Stack.Screen
            options={{
                title: 'Hello2'
            }}
            ></Stack.Screen>
            <Text>Home page</Text>
            {/* <Link onPress={()=>router.back()}>Go Back</Link> */}
        </View>
    );
  }