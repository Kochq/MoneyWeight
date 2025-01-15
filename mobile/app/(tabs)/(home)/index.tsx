import { useEffect, useState } from "react";
import { View, Text, StyleSheet, FlatList, Platform } from "react-native";
import TransactionCard from "./Transaction";
import { Transaction } from "@/types";

export default function HomeScreen() {
    const [transactions, setTransactions] = useState<Transaction[]>([]);

    const getTransactions = async () => {
        let url = "http://192.168.100.195:8080/api/transactions/details";

        const response = await fetch(url);
        const data = await response.json();

        console.log("Data fetched!");
        setTransactions(data.data);
    };

    useEffect(() => {
        console.log("");
        console.log("Fetching data...");
        getTransactions();
    }, []);

    return (
        <View style={styles.container}>
            <Text style={{ color: "#fff" }}>Home</Text>
            <FlatList
                style={styles.list}
                contentContainerStyle={styles.listContent}
                data={transactions}
                renderItem={({ item }) => <TransactionCard t={item} />}
                keyExtractor={(item) => item.id.toString()}
            />
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#222",
        gap: 10,
    },

    list: {
        width: "100%",
    },

    link: {
        color: "blue",
        backgroundColor: "lightblue",
        padding: 10,
        borderRadius: 5,
    },

    listContent: {
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
        gap: 10,
    },
});
