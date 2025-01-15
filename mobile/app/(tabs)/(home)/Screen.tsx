import { View, StyleSheet } from "react-native";
import React from "react";

interface ScreenProps {
    children: React.ReactNode;
}

export default function Screen({ children }: ScreenProps) {
    return <View style={styles.container}>{children}</View>;
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#222",
        paddingTop: 16,
        paddingHorizontal: 12,
    },
});
