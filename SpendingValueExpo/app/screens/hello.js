import { Link, Stack, router } from 'expo-router';
import { Text, View, Pressable } from 'react-native';

export default function HelloPage() {
    return (
        <View>
            <Stack.Screen
            options={{

            }}
            ></Stack.Screen>
            <Text>Home page</Text>
            <Pressable onPress={()=> router.push('/screens/hello2')}>
                <Text>Go to 2</Text>
            </Pressable>
        </View>
    );
  }