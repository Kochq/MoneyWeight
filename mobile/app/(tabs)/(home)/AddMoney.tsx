import { View, TextInput, Button } from "react-native";
import Screen from "./Screen";
import { useEffect, useState } from "react";
import { Picker } from "@react-native-picker/picker";
import { Category } from "@/types";
import { useTheme } from "../../../theme/ThemeContext";

export default function AddMoney() {
    const { colors } = useTheme();
    const [name, setName] = useState<string>("");
    const [money, setMoney] = useState("");
    const [categories, setCategories] = useState<Category[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<Category>();

    const addTransaction = async () => {
        let url = "http://serkq.org:8080/api/transactions";

        console.log("Name: " + name);
        console.log("Money: " + money);
        console.log("Category: " + selectedCategory);
        try {
            const response = await fetch(url, {
                method: "POST",
                body: JSON.stringify({
                    title: name,
                    amount: Number(money),
                    category_id: 1,
                    subcategory_id: 1,
                    currency: "ARS",
                    payment_method: "cosa",
                    exchange_rate: 1500,
                    notes: "esto es una nota",
                    from_account_id: 1,
                }),
            });

            const data = await response.json();

            console.log("Transaction added!");
            console.log(data);
        } catch (e) {
            console.error(e);
        }
    };

    const getCategories = async () => {
        let url = "http://serkq.org:8080/api/categories";

        const response = await fetch(url);
        const data = await response.json();

        console.log("Data fetched!");
        setCategories(data.data);
    };

    useEffect(() => {
        console.log("");
        console.log("Fetching data...");
        getCategories();
    }, []);

    return (
        <Screen>
            <View
                style={{
                    width: "100%",
                }}
            >
                <TextInput
                    placeholder="Name"
                    value={name}
                    onChangeText={setName}
                />
                <TextInput
                    placeholder="Money"
                    keyboardType="numeric"
                    value={money}
                    onChangeText={setMoney}
                />
                <Picker
                    style={{ borderWidth: 1, color: colors.textPrimary }}
                    selectedValue={selectedCategory}
                    onValueChange={(itemValue, itemIndex) =>
                        setSelectedCategory(itemValue)
                    }
                >
                    {categories.map((category) => (
                        <Picker.Item
                            key={category.id}
                            label={category.name}
                            value={category.id}
                        />
                    ))}
                </Picker>
                <Button title="Add" onPress={addTransaction} />
            </View>
        </Screen>
    );
}
