import { useEffect, useState } from "react";
import { StyleSheet, FlatList, RefreshControl } from "react-native";
import TransactionCard from "@/components/Transaction";
import TotalCard from "@/components/TotalCard";
import Buttons from "@/components/Buttons";
import { Transaction } from "@/types";
import Screen from "./Screen";

export default function HomeScreen() {
    const [transactions, setTransactions] = useState<Transaction[]>([]);
    const [refreshing, setRefreshing] = useState(false);

    const onRefresh = () => {
        setRefreshing(true);
        getTransactions();
        setRefreshing(false);
    };

    const getTransactions = async () => {
        let url = "http://serkq.org:8080/api/transactions/details";

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
                refreshControl={
                    <RefreshControl
                        refreshing={refreshing}
                        onRefresh={onRefresh}
                        colors={["grey"]}
                        progressBackgroundColor={"black"}
                    />
                }
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
