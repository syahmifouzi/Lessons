import { useContext } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import * as eva from '@eva-design/eva';
import { ApplicationProvider, IconRegistry } from '@ui-kitten/components';
import { EvaIconsPack } from '@ui-kitten/eva-icons';
import { MainNavigator } from './app/navigations/MainNavigator';
import {
  SafeAreaProvider,
} from 'react-native-safe-area-context';
import { MobxContext, ContextStores } from './app/stores/Context';


// Optionally import the services that you want to use
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries
// import {...} from "firebase/auth";
// import {...} from "firebase/database";
// import {...} from "firebase/firestore";
// import {...} from "firebase/functions";
// import {...} from "firebase/storage";
// import { getAnalytics } from "firebase/analytics";

export default function App() {
  const mobxCtx = useContext(MobxContext);
  return (
    <>
      <IconRegistry icons={EvaIconsPack} />
      <ApplicationProvider {...eva} theme={eva.light}>
        <SafeAreaProvider>
          <NavigationContainer>
            <MobxContext.Provider value={mobxCtx}>
              <MainNavigator />
            </MobxContext.Provider>
          </NavigationContainer>
        </SafeAreaProvider>
      </ApplicationProvider>
    </>
  );
}
