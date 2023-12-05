import { Tabs } from 'expo-router';
import Ionicons from '@expo/vector-icons/Ionicons';

export default function Layout() {
  return (
    <Tabs>
        <Tabs.Screen 
        name='home'
        options={{
            title: 'Tab one',
            tabBarIcon: ({focused, color, size}) => {

                return <Ionicons name={'ios-information-circle'} size={size} color={color} />;
            }
        }} />
        <Tabs.Screen 
        name='home4'
        options={{
            title: 'Tab two',
            tabBarIcon: ({focused, color, size}) => {

                return <Ionicons name={'ios-information-circle'} size={size} color={color} />;
            }
        }} />
    </Tabs>
  );
}