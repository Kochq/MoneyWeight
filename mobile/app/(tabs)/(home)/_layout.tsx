import { Stack } from "expo-router";
import React from "react";
import AccountSelector from "../../../components/AccountSelector";

export default function HomeLayout() {
    return (
        <Stack
            screenOptions={{
                headerShown: true,
                headerStyle: {
                    backgroundColor: "#121212",
                },
                headerTintColor: "#fff",
                headerTitleStyle: {
                    fontWeight: "bold",
                },
            }}
        >
            <Stack.Screen
                name="index"
                options={{
                    headerTitle: () => <AccountSelector />,
                    headerTitleAlign: "left",
                }}
            />
            <Stack.Screen name="details/[id]" />
            <Stack.Screen name="transaction/[id]" />
            <Stack.Screen name="AddMoney" />
        </Stack>
    );
}
