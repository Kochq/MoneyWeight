import { useEffect, useState } from "react";
import { StyleSheet, FlatList } from "react-native";
import TransactionCard from "./Transaction";
import { Transaction } from "@/types";
import { TotalCard } from "./TotalCard";
import Screen from "./Screen";
import { Buttons } from "./Buttons";

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
        <Screen>
            <TotalCard />
            <Buttons />
            <FlatList
                style={styles.list}
                contentContainerStyle={styles.listContent}
                data={transactions}
                renderItem={({ item }) => <TransactionCard t={item} />}
                keyExtractor={(item) => item.id.toString()}
            />
        </Screen>
    );
}

const styles = StyleSheet.create({
    list: {
        width: "100%",
    },

    listContent: {
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
    },
});
