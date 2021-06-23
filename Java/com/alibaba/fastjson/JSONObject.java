import com.alibaba.fastjson.JSONObject;
import com.alibaba.fastjson.Feature;

public class {
	public static void main() {
		// 解析text json文本， 得到JSONObject对象
		JSONObject respondeBodyJson = JSONObject.parseObject(str);

		// 解析为有序对象(JSONObjectm默认为有序对象, 生成json字符串会进行排序)
		JSONObject respondeBodyJson = JSONObject.parseObject(str, Feature.OrderedField);

		// 创建JSONObject有序对象
		JSONObject jso = new JSONObject(true);

		// 添加对象
		jso.put("id", "123");

		// 转换为json
		String text = JSONObject.toJsonString(jso);
	}
}
