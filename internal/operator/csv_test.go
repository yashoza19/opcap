package operator

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/opdev/opcap/internal/logger"
	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("CSV", func() {
	csv := operatorv1alpha1.ClusterServiceVersion{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testcsv",
			Namespace: "testns",
		},
	}

	logger.InitLogger("debug")
	scheme := runtime.NewScheme()
	operatorv1alpha1.AddToScheme(scheme)
	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(&csv).Build()
	var operatorClient operatorClient = operatorClient{
		Client: client,
	}
	Context("Test CSVs", func() {
		It("should exercise CSVs", func() {
			By("getting a CSV", func() {
				csv, err := operatorClient.GetCompletedCsvWithTimeout(context.TODO(), "testns", time.Minute)
				Expect(err).ToNot(HaveOccurred())
				Expect(csv).ToNot(BeNil())
			})
		})
	})
})
