import { useLocalSearchParams } from "expo-router";
import { View, Text, StyleSheet, Pressable } from "react-native";

export default function TransactionScreen() {
    const { id } = useLocalSearchParams();

    const removeTransaction = async (id: any) => {
        let url = `http://serkq.org:8080/api/transactions/${id}`;

        // formatDate: "2023-12-27 20:18:43",
        try {
            const response = await fetch(url, { method: "DELETE" });

            const data = await response.json();

            console.log("Transaction deleted!");
            alert(data.status);
        } catch (e) {
            console.error(e);
        }
    };

    return (
        <View style={styles.container}>
            <Text>Details of transaction {id} </Text>
            <Pressable
                onPress={() => {
                    removeTransaction(id);
                }}
            >
                <Text>Delete transaction</Text>
            </Pressable>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
});
