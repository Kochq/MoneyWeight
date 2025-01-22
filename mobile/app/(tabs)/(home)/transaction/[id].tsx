import { useLocalSearchParams } from "expo-router";
import { View, Text, StyleSheet, Pressable } from "react-native";

export default function TransactionScreen() {
    const { id } = useLocalSearchParams();
    const apiUrl = process.env.EXPO_PUBLIC_API_BASE;

    const removeTransaction = async (id: any) => {
        const url = apiUrl + `/api/transactions/${id}`;

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
