import { Link } from "expo-router";
import React from "react";
import { Pressable, StyleSheet, Text, View } from "react-native";

export default function Settings() {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Settings</Text>
            <Text style={styles.text}>This is the settings page.</Text>
            <Link href="/details/1" asChild>
                <Pressable style={styles.link}>
                    <Text>Go to first details</Text>
                </Pressable>
            </Link>
            <Link href="/details/2" asChild>
                <Pressable style={styles.link}>
                    <Text>Go to second details</Text>
                </Pressable>
            </Link>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#222",
    },

    link: {
        color: "blue",
        backgroundColor: "lightblue",
        padding: 10,
        borderRadius: 5,
        marginBottom: 10,
    },

    title: {
        color: "#fff",
        fontSize: 20,
        fontWeight: "bold",
    },

    text: {
        color: "#fff",
        fontSize: 16,
        marginBottom: 10,
    },
});
