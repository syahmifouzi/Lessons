import { View } from 'react-native';

import { Link } from 'expo-router';


export default function Page() {
  return (
    <View>
      <Link href="/mytabs/home4">About</Link>

      {/* <Link href="/user/bacon">View user</Link> */}
    </View>
  );
}
